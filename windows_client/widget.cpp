#include "widget.h"
#include <QDebug>

Widget::Widget(QWidget *parent)
    : QWidget(parent)
{
    _WIDTH = static_cast<int>(GetSystemMetrics(SM_CXSCREEN)/2);
    _HEIGHT = static_cast<int>(GetSystemMetrics(SM_CYSCREEN)/2);

    setWindowTitle(tr("projectZ for windows"));
    setMinimumSize(_WIDTH,_HEIGHT);
    setMaximumSize(_WIDTH,_HEIGHT);
    setStyleSheet("color: gray; background: silver;");//border: 1px solid green;

    std::string address("*.zgg2001.com:8888");
    auto channel = grpc::CreateChannel(address, grpc::InsecureChannelCredentials());
    _stub = ProjectService::NewStub(channel);

    ClientContext context;
    UserLoginRequest request;
    UserLoginResponse response;

    request.set_username("test");
    request.set_password("test123456");

    Status status = _stub->UserLogin(&context, request, &response);
    if(status.ok())
    {
        qDebug()<<"success";
    }
    else
    {
        qDebug()<< status.error_code();
        qDebug()<< QString::fromStdString(status.error_message());
    }
}

Widget::~Widget()
{
}

std::string Widget::read_file(std::string file_path)
{
    std::ifstream f{file_path.c_str(), std::ios::binary};
    std::stringstream buffer;
    buffer << f.rdbuf();
    f.close();
    return buffer.str();
}




