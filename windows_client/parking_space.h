#ifndef PARKING_SPACE_H
#define PARKING_SPACE_H

#include <string>

class parking_space
{
public:
    enum Alarm {ALARM_NO, ALARM_FIRE, ALARM_GAS, ALARM_FIRE_AND_GAS};

public:
    parking_space(int id);

private:
    bool _use = false;
    int _id = 0;
    int _temperature = 0;
    int _humidity = 0;
    Alarm _alarm = ALARM_NO;
    long long _entryTime = 0;
    std::string _license = "";
};

#endif // PARKING_SPACE_H
