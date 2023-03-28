#include "my_mqtt.h"

mqtt_client::mqtt_client(int id, const char* ip, uint16_t port, const char* name, const char* passwd, const char* pub, const char* sub)
  : _esp32_id(id),
    _server_ip(ip),
    _server_port(port),
    _username(name),
    _password(passwd),
    _pub_topic(pub),
    _sub_topic(sub) {
}

mqtt_client::~mqtt_client() {
  delete _sub;
  delete _pub;
  delete _mqtt;
}

void mqtt_client::init() {
  _mqtt = new Adafruit_MQTT_Client(&_client, _server_ip, _server_port, _username, _password);
  _pub = new Adafruit_MQTT_Publish(_mqtt, _pub_topic);
  _sub = new Adafruit_MQTT_Subscribe(_mqtt, _sub_topic);
  _mqtt->subscribe(_sub);
}

void mqtt_client::mqtt_connect() {
  int8_t ret;
  if (_mqtt->connected()) {
    return;
  }
  while ((ret = _mqtt->connect()) != 0) {
    _mqtt->disconnect();
    vTaskDelay(pdMS_TO_TICKS(5000));
  }
}

void mqtt_client::mqtt_pub(const string& msg) {
  mqtt_connect();
  _pub->publish(msg.c_str());
}

int mqtt_client::mqtt_sub(int millisecond, string& msg) {
  mqtt_connect();
  Adafruit_MQTT_Subscribe* subscription;
  while ((subscription = _mqtt->readSubscription(millisecond))) {
    if (subscription == _sub) {
      msg = (char*)_sub->lastread;
      return HAVE_NEW_MSG;
    }
  }
  return NO_NEW_MSG;
}

int mqtt_client::parse_cmd(const string& msg) {
  size_t pos = msg.find(":");
  string id_str = msg.substr(0, pos);
  string cmd_str = msg.substr(pos + 1, msg.size());
  int id, cmd;
  try {
    id = stoi(id_str);
    cmd = stoi(cmd_str);
  } catch (...) {
    return INVALID_CMD;
  }
  if (id != _esp32_id) {
    return INVALID_CMD;
  }
  return cmd;
}
