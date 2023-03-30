#include "widget.h"
#include <QDebug>

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
        l = new QLabel(tr("豫A66666"));//文本提示
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
}

Widget::~Widget()
{
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

void Widget::rpc_test()
{
    std::string address("*.zgg2001.com:11110");
    auto channel = grpc::CreateChannel(address, grpc::InsecureChannelCredentials());
    _stub = ProjectService::NewStub(channel);

    ClientContext context;
    AdminLoginRequest request;
    AdminLoginResponse response;

    request.set_p_id(2);
    request.set_password("75121");

    Status status = _stub->AdminLogin(&context, request, &response);
    if(status.ok())
    {
        qDebug() << "success" << response.count() << response.result();
    }
    else
    {
        qDebug()<< status.error_code();
        qDebug()<< QString::fromStdString(status.error_message());
    }
}

std::string Widget::read_file(std::string file_path)
{
    std::ifstream f{file_path.c_str(), std::ios::binary};
    std::stringstream buffer;
    buffer << f.rdbuf();
    f.close();
    return buffer.str();
}




