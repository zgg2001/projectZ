#include "widget.h"
#include <QDebug>

using namespace std;

void delivered(void *context, MQTTClient_deliveryToken dt);
int msgarrvd(void *context, char *topicName, int topicLen, MQTTClient_message *message);
void connlost(void *context, char *cause);

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
    int rc;
    if ((rc = MQTTClient_disconnect(_mqtt_client, 10000)) != MQTTCLIENT_SUCCESS)
    {
        printf("Failed to disconnect, return code %d\n", rc);
        rc = EXIT_FAILURE;
    }
    MQTTClient_destroy(&_mqtt_client);
}

bool Widget::mqtt_connect(string ip)
{
    string addr = ip + MQTT_PORT;
    MQTTClient_connectOptions conn_opts = MQTTClient_connectOptions_initializer;
    int rc;
    if ((rc = MQTTClient_create(&_mqtt_client, addr.c_str(), "windows_client",
            MQTTCLIENT_PERSISTENCE_NONE, nullptr)) != MQTTCLIENT_SUCCESS)
    {
        printf("Failed to create client, return code %d\n", rc);
        rc = EXIT_FAILURE;
        return false;
    }
    if ((rc = MQTTClient_setCallbacks(_mqtt_client, nullptr, connlost, msgarrvd, delivered)) != MQTTCLIENT_SUCCESS)
    {
        printf("Failed to set callbacks, return code %d\n", rc);
        rc = EXIT_FAILURE;
        return false;
    }
    conn_opts.keepAliveInterval = 20;
    conn_opts.cleansession = 1;
    conn_opts.username = MQTT_USERNAME;
    conn_opts.password = MQTT_PASSWORD;
    if ((rc = MQTTClient_connect(_mqtt_client, &conn_opts)) != MQTTCLIENT_SUCCESS)
    {
        printf("Failed to connect, return code %d\n", rc);
        rc = EXIT_FAILURE;
        return false;
    }
    if ((rc = MQTTClient_subscribe(_mqtt_client, Widget::SUB_TOPIC, Widget::MQTT_QOS)) != MQTTCLIENT_SUCCESS)
    {
        printf("Failed to subscribe, return code %d\n", rc);
        rc = EXIT_FAILURE;
        return false;
    }
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

void delivered(void *context, MQTTClient_deliveryToken dt)
{
    printf("Message with token value %d delivery confirmed\n", dt);
    //deliveredtoken = dt;
}

int msgarrvd(void *context, char *topicName, int topicLen, MQTTClient_message *message)
{
    qDebug() << "Message arrived";
    qDebug() << "topic: "<< topicName;
    qDebug() << "message: " << message->payloadlen << (char*)message->payload;
    MQTTClient_freeMessage(&message);
    MQTTClient_free(topicName);
    return 1;
}

void connlost(void *context, char *cause)
{
    qDebug() << "Connection lost";
    qDebug() << "cause: " << cause;
}



