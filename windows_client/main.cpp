#include "dialog.h"


#include <QTextCodec>
#include <QApplication>

int main(int argc, char *argv[])
{
    QTextCodec *codec = QTextCodec::codecForName("UTF-8");
    QTextCodec::setCodecForLocale(codec);
    QApplication a(argc, argv);
    // login
    Dialog d;
    d.show();
    return a.exec();
}
