#include "parking_space.h"

parking_space::parking_space(int id):
    _id(id)
{

}

void parking_space::set_license_and_entrytime(bool use, std::string license, long long entrytime)
{
    _use = use;
    _license = license;
    _entrytime = entrytime;
}
