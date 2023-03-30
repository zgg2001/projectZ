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

#include <grpcpp/grpcpp.h>
#include "pbfile/service.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

class Widget : public QWidget
{
    Q_OBJECT

public:
    Widget(QWidget *parent = nullptr);
    ~Widget();

private:
    std::string read_file(std::string file_path);
    void Widget::rpc_test();

    void paintEvent(QPaintEvent *);//重写函数

private:
    int _WIDTH;
    int _HEIGHT;

    std::unique_ptr<ProjectService::Stub> _stub;

    std::vector<QPushButton*> _parking_space_buttons;
    std::vector<QLabel*> _parking_space_labels;
};
#endif // WIDGET_H
