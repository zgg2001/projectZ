#ifndef PARKING_SPACE_H
#define PARKING_SPACE_H

#include <string>
#include <QString>

class parking_space
{
public:
    enum Alarm {ALARM_NO, ALARM_FIRE, ALARM_GAS, ALARM_FIRE_AND_GAS};

public:
    parking_space(int id);
    void set_license_and_entrytime(bool use, std::string license, QString entrytime);
    void set_data(int t, int h, Alarm a);
    bool is_use() { return _use; }
    int get_id() { return _id; }
    int get_temp() { return _temperature; }
    int get_humi() { return _humidity; }
    Alarm get_alarm() { return _alarm; }
    QString get_entrytime() { return _entrytime; }
    std::string get_license() { return _license; }

private:
    bool _use = false;
    int _id = 0;
    int _temperature = 0;
    int _humidity = 0;
    Alarm _alarm = ALARM_NO;
    QString _entrytime = 0;
    std::string _license = "";
};

#endif // PARKING_SPACE_H
