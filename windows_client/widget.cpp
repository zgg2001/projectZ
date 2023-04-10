#include "widget.h"
#include <QDebug>

using namespace std;

bool Widget::_mqtt_subscribed = false;
bool Widget::_mqtt_finished = false;

void connlost(void *context, char *cause);
int msgarrvd(void *context, char *topicName, int topicLen, MQTTAsync_message *message);
void onConnect(void* context, MQTTAsync_successData* response);
void onConnectFailure(void* context, MQTTAsync_failureData* response);
void onSubscribe(void* context, MQTTAsync_successData* response);
void onSubscribeFailure(void* context, MQTTAsync_failureData* response);

Widget::Widget(QWidget *parent)
    : QWidget(parent)
{
    _WIDTH = static_cast<int>(GetSystemMetrics(SM_CXSCREEN)*2/3);
    _HEIGHT = static_cast<int>(GetSystemMetrics(SM_CYSCREEN)*2/3);
    setWindowTitle(tr("projectZ for windows"));
    setMinimumSize(_WIDTH,_HEIGHT);
    setMaximumSize(_WIDTH,_HEIGHT);
    QFont font("微软雅黑", 15, QFont::Bold, false);//生成字体
    int buttonWidth = _WIDTH / 10;
    int buttonHeight = _HEIGHT / 10;
    // 车位出库按钮
    int x_interval = _WIDTH / 4, y_interval = _HEIGHT / 3;
    int x = 0, y = y_interval, count = 0;
    _parking_space_buttons.assign(8, nullptr);
    for(auto b : _parking_space_buttons) {
        b = new QPushButton(tr("出库"));
        b->setParent(this);
        b->move(x,y);
        b->setFont(font);
        b->setFixedSize(buttonWidth, buttonHeight);
        b->setFlat(true);
        ++count;
        x = x + x_interval;
        if (count == 4)
        {
            x = 0;
            y += y_interval;
        }
    }
    // 车牌号显示
    x = buttonWidth;
    y = y_interval;
    count = 0;
    _parking_space_labels.assign(8, nullptr);
    for(auto l : _parking_space_labels) {
        l = new QLabel(tr("-"));//文本提示
        l->setParent(this);
        l->move(x,y);
        l->setFont(font);
        l->setFixedSize(x_interval - buttonWidth, buttonHeight);
        ++count;
        x = x + x_interval;
        if (count == 4)
        {
            x = buttonWidth;
            y += y_interval;
        }
    }
    // 车位信息显示
    QFont font_info("微软雅黑", 10, false, false);
    x = buttonWidth / 2;
    y = y_interval + buttonHeight;
    count = 0;
    _parking_space_info_labels.assign(8, nullptr);
    for(auto l : _parking_space_info_labels) {
        l = new QLabel(tr(""));
        l->setParent(this);
        l->move(x,y);
        l->setFont(font_info);
        l->setFixedSize(x_interval, y_interval - buttonHeight);
        ++count;
        x = x + x_interval;
        if (count == 4)
        {
            x = buttonWidth / 2;
            y += y_interval;
        }
    }
    // rpc connect
    std::string address("*.zgg2001.com:11110");
    auto channel = grpc::CreateChannel(address, grpc::InsecureChannelCredentials());
    _stub = ProjectService::NewStub(channel);
}

Widget::~Widget()
{
    MQTTAsync_disconnectOptions disc_opts = MQTTAsync_disconnectOptions_initializer;
    MQTTAsync_destroy(&_mqtt_client);
}

bool Widget::mqtt_connect(string ip)
{
    string addr = ip + MQTT_PORT;
    MQTTAsync_connectOptions  conn_opts = MQTTAsync_connectOptions_initializer;
    int rc;

    _mqtt_subscribed = false;
    _mqtt_finished = false;
    if ((rc = MQTTAsync_create(&_mqtt_client, addr.c_str(), "windows_client",
            MQTTCLIENT_PERSISTENCE_NONE, nullptr)) != MQTTCLIENT_SUCCESS)
    {
        Widget::_mqtt_finished = true;
        return false;
    }
    if ((rc = MQTTAsync_setCallbacks(_mqtt_client, _mqtt_client, connlost, msgarrvd, nullptr)) != MQTTASYNC_SUCCESS)
    {
        Widget::_mqtt_finished = true;
        return false;
    }

    conn_opts.keepAliveInterval = 20;
    conn_opts.cleansession = 1;
    conn_opts.onSuccess = onConnect;
    conn_opts.onFailure = onConnectFailure;
    conn_opts.username = MQTT_USERNAME;
    conn_opts.password = MQTT_PASSWORD;
    if ((rc = MQTTAsync_connect(_mqtt_client, &conn_opts)) != MQTTCLIENT_SUCCESS)
    {
        Widget::_mqtt_finished = true;
        return false;
    }

    while(!_mqtt_subscribed && !_mqtt_finished)
        Sleep(100);
    if(_mqtt_finished)
        return false;
    return true;
}

void Widget::init_parking()
{
    _spaces.clear();
    for (int id = 1; id <= _space_count; ++id) {
        parking_space temp(id);
        string license = "";
        long long entrytime = 0;
        bool ok = rpc_get_space_info(_pid, id, license, entrytime);
        if(ok)
            temp.set_license_and_entrytime(true, license, entrytime);
        _spaces.push_back(std::move(temp));
    }
}

void Widget::paintEvent(QPaintEvent *)
{
    QPainter painter(this);//建立QPainter并绑定widget
    QPen pen;//生成画笔
    QFont font1("微软雅黑",15,QFont::Bold,true);//生成字体
    pen.setColor(Qt::gray);//笔 灰色
    pen.setStyle(Qt::SolidLine);
    pen.setWidthF(2);//此函数可定义宽度可精确到小数
    painter.setPen(pen);//画笔画布绑定
    painter.setFont(font1);//字体画布绑定
    painter.setViewport(0, 0, _WIDTH, _HEIGHT);//设置画布视窗大小大小为800*800大小
    painter.setWindow(0, 0, _WIDTH, _HEIGHT);//设置逻辑坐标范围
    painter.fillRect(0, 0, _WIDTH, _HEIGHT/3, Qt::white);
    painter.fillRect(0, _HEIGHT/3, _WIDTH, _HEIGHT, Qt::transparent);
    painter.setRenderHint(QPainter::Antialiasing,true);//抗锯齿
    int x_interval = _WIDTH / 4, y_interval = _HEIGHT / 3 * 2, x_i = _WIDTH / 128, y_i = _HEIGHT / 96;
    for(int now = 1; now <= 3; ++now)
    {
        int x = x_interval * now;
        painter.drawLine(x - x_i, _HEIGHT/3*2, x + x_i, _HEIGHT/3*2);
        painter.drawLine(x, y_interval - y_i, x, y_interval + y_i);
    }
}

bool Widget::rpc_get_space_info(int pid, int sid, string& license, long long& entrytime)
{
    ClientContext context;
    AdminGetSpaceInfoRequest request;
    AdminGetSpaceInfoResponse response;

    request.set_p_id(pid);
    request.set_s_id(sid);

    Status status = _stub->AdminGetSpaceInfo(&context, request, &response);
    if(status.ok())
    {
        license = response.license();
        entrytime = response.entrytime();
        if(!response.is_use())
            return false;
        return true;
    }
    return false;
}

void connlost(void *context, char *cause)
{
    MQTTAsync client = static_cast<MQTTAsync>(context);
    MQTTAsync_connectOptions conn_opts = MQTTAsync_connectOptions_initializer;
    int rc;
    conn_opts.keepAliveInterval = 20;
    conn_opts.cleansession = 1;
    conn_opts.onSuccess = onConnect;
    conn_opts.onFailure = onConnectFailure;
    conn_opts.username = Widget::MQTT_USERNAME;
    conn_opts.password = Widget::MQTT_PASSWORD;
    if ((rc = MQTTAsync_connect(client, &conn_opts)) != MQTTASYNC_SUCCESS)
    {
        Widget::_mqtt_finished = true;
    }
}

int msgarrvd(void *context, char *topicName, int topicLen, MQTTAsync_message *message)
{
    qDebug() << "Message arrived";
    qDebug() << "topic: " << topicName;
    qDebug() << "message:" << message->payloadlen << static_cast<char*>(message->payload);
    MQTTAsync_freeMessage(&message);
    MQTTAsync_free(topicName);
    return 1;
}

void onConnect(void* context, MQTTAsync_successData* response)
{
    MQTTAsync client = static_cast<MQTTAsync>(context);
    MQTTAsync_responseOptions opts = MQTTAsync_responseOptions_initializer;
    int rc;
    opts.onSuccess = onSubscribe;
    opts.onFailure = onSubscribeFailure;
    opts.context = client;
    if ((rc = MQTTAsync_subscribe(client, Widget::SUB_TOPIC, Widget::MQTT_QOS, &opts)) != MQTTASYNC_SUCCESS)
    {
        Widget::_mqtt_finished = true;
    }
}

void onConnectFailure(void* context, MQTTAsync_failureData* response)
{
    Widget::_mqtt_finished = true;
}

void onSubscribe(void* context, MQTTAsync_successData* response)
{
    Widget::_mqtt_subscribed = true;
}

void onSubscribeFailure(void* context, MQTTAsync_failureData* response)
{
    Widget::_mqtt_finished = true;
}



