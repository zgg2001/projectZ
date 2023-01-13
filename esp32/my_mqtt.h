#ifndef _MY_MQTT_H_
#define _MY_MQTT_H_

#include <string>
#include "WiFiClient.h"
#include "Adafruit_MQTT.h"
#include "Adafruit_MQTT_Client.h"

using std::string;

class mqtt_client {
public:
  //base
  mqtt_client(const char* ip, uint16_t port, const char* name, const char* passwd, const char* pub, const char* sub);
  mqtt_client() = delete;
  mqtt_client(const mqtt_client&) = delete;
  mqtt_client& operator=(const mqtt_client&) = delete;
  virtual ~mqtt_client();

  //
  void init();
  void mqtt_connect();
  void mqtt_pub(string msg);
  void mqtt_sub(int millisecond);

private:
  WiFiClient _client;
  const char* _server_ip;
  uint16_t _server_port;
  const char* _username;
  const char* _password;
  const char* _pub_topic;
  const char* _sub_topic;
  Adafruit_MQTT_Client* _mqtt;
  Adafruit_MQTT_Publish* _pub;
  Adafruit_MQTT_Subscribe* _sub;
};

#endif
