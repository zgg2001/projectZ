#ifndef DIALOG_H
#define DIALOG_H

#include "widget.h"

#include <QDialog>
#include <QCloseEvent>
#include <QMessageBox>

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
    bool rpc_login(int pid, std::string password, int& count, LoginResult& result);

    void closeEvent(QCloseEvent *event);

private slots:
    void on_pushButton_2_clicked();

private:
    Ui::Dialog *ui;
    Widget _w;
    // rpc
    std::unique_ptr<ProjectService::Stub> _stub;
};

#endif // DIALOG_H
