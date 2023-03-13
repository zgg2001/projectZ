#ifndef WIDGET_H
#define WIDGET_H

#include <QWidget>
#include <windows.h>

class Widget : public QWidget
{
    Q_OBJECT

public:
    Widget(QWidget *parent = nullptr);
    ~Widget();

private:
    int _WIDTH;
    int _HEIGHT;

};
#endif // WIDGET_H
