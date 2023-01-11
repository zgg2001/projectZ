#!/usr/bin/python3
# -*- coding: UTF-8 -*-

import RPi.GPIO as GPIO
import time
from hyperlpr import *
import cv2
import os

GPIO.setmode(GPIO.BCM)

# 摄像头超声波传感
Trig_Pin1 = 22
Echo_Pin1 = 27
GPIO.setup(Trig_Pin1, GPIO.OUT, initial = GPIO.LOW)
GPIO.setup(Echo_Pin1, GPIO.IN)

# 挡杆超声波传感
Trig_Pin2 = 6
Echo_Pin2 = 5
GPIO.setup(Trig_Pin2, GPIO.OUT, initial = GPIO.LOW)
GPIO.setup(Echo_Pin2, GPIO.IN)

# 舵机引脚/舵机状态机标识/角度
servo_pin = 17
servo_status = 0 # 0-关闭 1-验证 2-待开启 3-已开启
servo_lift = 90
servo_down = 0
GPIO.setup(servo_pin, GPIO.OUT)

# 车牌/车牌识别最短间隔(s)/抬杠最短时间(s)
plate = "豫A88888"
detection_interval = 2
rise_time = 10

# 读取摄像头设备/设置分辨率/创建tmp目录
cap = cv2.VideoCapture('/dev/video0', cv2.CAP_V4L)
cap.set(cv2.CAP_PROP_FRAME_WIDTH, 1280)
cap.set(cv2.CAP_PROP_FRAME_HEIGHT, 720)
cap.set(cv2.CAP_PROP_BUFFERSIZE, 1)
dirs = './tmp'
if not os.path.exists(dirs):
    os.makedirs(dirs)

# 管道
pipe_read  = "./tmp/pipe.2"
pipe_write = "./tmp/pipe.1"
rf = os.open(pipe_read, os.O_RDONLY | os.O_NONBLOCK)
wf = os.open(pipe_write, os.O_SYNC | os.O_CREAT | os.O_RDWR)

time.sleep(1)

# get distance
def checkdist(trig, echo):
    GPIO.output(trig, GPIO.HIGH)
    time.sleep(0.00001)
    GPIO.output(trig, GPIO.LOW)
    while GPIO.input(echo) == 0:
        pass
    t1 = time.time()
    while GPIO.input(echo) == 1:
        pass
    t2 = time.time()
    return (t2 - t1) * 340 * 100 / 2

# set servo
def setServoAngle(servo, angle):
	pwm = GPIO.PWM(servo, 50)
	pwm.start(8)
	dutyCycle = angle / 18. + 3.
	pwm.ChangeDutyCycle(dutyCycle)
	time.sleep(0.3)
	pwm.stop()

if __name__ == '__main__':
    try:
        last_time = time.time()
        while True:
            ret, frame = cap.read()
            if servo_status == 0 and time.time() - last_time > detection_interval:
                distance1 = checkdist(Trig_Pin1, Echo_Pin1)
                # print(distance1)
                if distance1 < 30:
                    # 拍照
                    cv2.imwrite('./tmp/tmp.jpg', frame) 
                    # 分析
                    image = cv2.imread('./tmp/tmp.jpg')
                    res = HyperLPR_plate_recognition(image)
                    if len(res):
                        plate = res[0][0]
                        msg = plate.encode() + bytes('\n', 'utf-8')
                        os.write(wf, msg)
                        servo_status = 1
                    last_time = time.time()
                else:
                    time.sleep(1)
            elif servo_status == 1:
                # 验证逻辑
                # print(plate)
                ret = os.read(rf, 10).decode()
                # print(ret)
                if ret == "pass":
                    servo_status = 2
                else:
                    servo_status = 0
            elif servo_status == 2:
                # 抬杆
                # print("lift")
                last_time = time.time()
                setServoAngle(servo_pin, servo_lift)
                servo_status = 3
            elif servo_status == 3 and time.time() - last_time > rise_time:
                distance2 = checkdist(Trig_Pin2, Echo_Pin2)
                # print(distance2)
                if distance2 > 10:
                    # 落杆
                    # print("down")
                    setServoAngle(servo_pin, servo_down)
                    servo_status = 0
                else:
                    time.sleep(1)
    except BaseException:
        GPIO.cleanup()
        cap.release()
