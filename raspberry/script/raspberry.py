import RPi.GPIO as GPIO
import time
from hyperlpr import *
import cv2
import os

# 摄像头超声波传感
Trig_Pin1 = 22
Echo_Pin1 = 27
GPIO.setmode(GPIO.BCM)
GPIO.setup(Trig_Pin1, GPIO.OUT, initial = GPIO.LOW)
GPIO.setup(Echo_Pin1, GPIO.IN)

# 车牌识别最短间隔/秒
detection_interval = 2
# 读取摄像头设备/设置分辨率/创建tmp目录
cap = cv2.VideoCapture('/dev/video0') #, cv2.CAP_V4L
cap.set(cv2.CAP_PROP_FRAME_WIDTH, 1280)
cap.set(cv2.CAP_PROP_FRAME_HEIGHT, 720)
dirs = './tmp'
if not os.path.exists(dirs):
    os.makedirs(dirs)

time.sleep(1)

def checkdist():
    GPIO.output(Trig_Pin1, GPIO.HIGH)
    time.sleep(0.00001)
    GPIO.output(Trig_Pin1, GPIO.LOW)
    while GPIO.input(Echo_Pin1) == 0:
        pass
    t1 = time.time()
    while GPIO.input(Echo_Pin1) == 1:
        pass
    t2 = time.time()
    return (t2 - t1) * 340 * 100 / 2

try:
    last_time = time.time()
    while True:
        ret, frame = cap.read()
        distance = checkdist()
        if distance < 25 and time.time() - last_time > detection_interval:
            # 拍照
            cv2.imwrite('./tmp/tmp.jpg', frame) 
            # 分析
            image = cv2.imread('./tmp/tmp.jpg')
            print(HyperLPR_plate_recognition(image))
            last_time = time.time()
except KeyboardInterrupt:
    GPIO.cleanup()
    cap.release()
    cv2.destroyAllWindows()
