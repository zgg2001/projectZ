#ifndef WIDGET_H
#define WIDGET_H

#include <windows.h>

#include <QWidget>
#include <QPushButton>
#include <QLabel>
#include <QPainter>

#include <string>
#include <fstream>
#include <sstream>
#include <utility>

#include <MQTTClient.h>

#include <grpcpp/grpcpp.h>
#include "pbfile/service.grpc.pb.h"

#include "parking_space.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

class Widget : public QWidget
{
    Q_OBJECT

private:
    static constexpr const char* MQTT_PORT = ":1883";
    static constexpr const char* MQTT_USERNAME = "test0";
    static constexpr const char* MQTT_PASSWORD = "z123456";

public:
    Widget(QWidget *parent = nullptr);
    ~Widget();
    void set_pid(int pid) { _pid = pid; }
    void set_parking_space_count(int count) { _space_count = count; }
    bool mqtt_connect(std::string ip);
    void init_parking();

private:
    void paintEvent(QPaintEvent *);//重写函数

private:
    int _WIDTH;
    int _HEIGHT;

    std::vector<QPushButton*> _parking_space_buttons;
    std::vector<QLabel*> _parking_space_labels;
    std::vector<QLabel*> _parking_space_info_labels;

    int _pid;
    int _space_count;
    std::vector<parking_space> _spaces;
    // rpc
    std::unique_ptr<ProjectService::Stub> _stub;
    // mqtt
    std::string _mqtt_ip;
    MQTTClient _mqtt_client;
};
#endif // WIDGET_H
