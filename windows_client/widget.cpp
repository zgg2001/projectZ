#include "widget.h"

Widget::Widget(QWidget *parent)
    : QWidget(parent)
{
    _WIDTH = static_cast<int>(GetSystemMetrics(SM_CXSCREEN)/2);
    _HEIGHT = static_cast<int>(GetSystemMetrics(SM_CYSCREEN)/2);

    setWindowTitle(tr("projectZ"));
    setMinimumSize(_WIDTH,_HEIGHT);
    setMaximumSize(_WIDTH,_HEIGHT);
    setStyleSheet("color: gray; background: silver;");//border: 1px solid green;
}

Widget::~Widget()
{
}

