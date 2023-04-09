#include "dialog.h"
#include "ui_dialog.h"
#include <QDebug>

using namespace std;

Dialog::Dialog(QWidget *parent) :
    QDialog(parent),
    ui(new Ui::Dialog)
{
    ui->setupUi(this);
    ui->lineEdit->setValidator(new QIntValidator(ui->lineEdit));
    ui->lineEdit_2->setValidator(new QRegExpValidator(QRegExp("[a-zA-Z0-9]+$")));
    ui->lineEdit_3->setValidator(new QRegExpValidator(QRegExp("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$")));
    // rpc connect
    std::string address("*.zgg2001.com:11110");
    auto channel = grpc::CreateChannel(address, grpc::InsecureChannelCredentials());
    _stub = ProjectService::NewStub(channel);
}

Dialog::~Dialog()
{
    delete ui;
}

void Dialog::on_pushButton_2_clicked()
{
    bool ok = false;
    int pid = ui->lineEdit->text().toInt();
    string password = ui->lineEdit_2->text().toStdString();
    string mqtt_ip = ui->lineEdit_3->text().toStdString();

    // rpc login
    int count = 0;
    LoginResult result = LOGIN_FAIL_NOT_EXIST;
    ok = rpc_login(pid, password, count, result);
    if(!ok)
        return;
    if(result == LOGIN_SUCCESS)
    {
        _w.set_pid(pid);
        _w.set_parking_space_count(count);
    }
    else if(result == LOGIN_FAIL_NOT_EXIST)
    {
        QMessageBox::information(nullptr, "Error", "User not exist!", QMessageBox::Yes, QMessageBox::Yes);
        return;
    }
    else if(result == LOGIN_FAIL_WRONG_PASSWORD)
    {
        QMessageBox::information(nullptr, "Error", "Wrong password!", QMessageBox::Yes, QMessageBox::Yes);
        return;
    }

    // mqtt connect
    ok = _w.mqtt_connect(mqtt_ip);
    if(!ok)
    {
        QMessageBox::information(nullptr, "Error", "Please check if the mqtt IP is correct!", QMessageBox::Yes, QMessageBox::Yes);
        return;
    }

    _w.init_parking();
    this->hide();
    _w.show();
}

bool Dialog::rpc_login(int pid, string password, int& count, LoginResult& result)
{
    ClientContext context;
    AdminLoginRequest request;
    AdminLoginResponse response;

    request.set_p_id(pid);
    request.set_password(password);

    Status status = _stub->AdminLogin(&context, request, &response);
    if(status.ok())
    {
        count = response.count();
        result = response.result();
        return true;
    }
    else
    {
        QMessageBox::information(nullptr, "Error", QString::fromStdString(status.error_message()), QMessageBox::Yes, QMessageBox::Yes);
        return false;
    }
}

void Dialog::closeEvent(QCloseEvent *event)
{
    event->accept();
}
