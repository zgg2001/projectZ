#ifndef DIALOG_H
#define DIALOG_H

#include "widget.h"

#include <QDialog>
#include <QCloseEvent>
#include <QMessageBox>

#include <MQTTClient.h>

#include <grpcpp/grpcpp.h>
#include "pbfile/service.grpc.pb.h"

namespace Ui {
class Dialog;
}

class Dialog : public QDialog
{
    Q_OBJECT

public:
    explicit Dialog(QWidget *parent = nullptr);
    ~Dialog();

private:
    bool mqtt_check(std::string ip);
    bool rpc_login(int pid, std::string password, int& count, LoginResult& result);

    void closeEvent(QCloseEvent *event);

private slots:
    void on_pushButton_2_clicked();

private:
    static constexpr const char* MQTT_PORT = ":1883";
    static constexpr const char* MQTT_USERNAME = "test0";
    static constexpr const char* MQTT_PASSWORD = "z123456";

private:
    Ui::Dialog *ui;
    Widget _w;
    std::unique_ptr<ProjectService::Stub> _stub;
};

#endif // DIALOG_H
