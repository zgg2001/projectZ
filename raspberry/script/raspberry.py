import RPi.GPIO as GPIO
import time
from hyperlpr import *
import cv2

# 摄像头超声波
Trig_Pin1 = 22
Echo_Pin1 = 27

GPIO.setmode(GPIO.BCM)
GPIO.setup(Trig_Pin1, GPIO.OUT, initial = GPIO.LOW)
GPIO.setup(Echo_Pin1, GPIO.IN)

# 读取摄像头设备/设置分辨率
cap = cv2.VideoCapture('/dev/video0', cv2.CAP_V4L)
cap.set(cv2.CAP_PROP_FRAME_WIDTH, 1280)
cap.set(cv2.CAP_PROP_FRAME_HEIGHT, 720)

time.sleep(2)

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
    while True:
        distance = checkdist()
        if distance < 25:
            time.sleep(1)
            ret, frame = cap.read()
            if ret:
                # 拍照
                print("ka cha")
                cv2.imwrite('./tmp/tmp1.jpg', frame) 
            # 分析
            image = cv2.imread("./tmp/tmp1.jpg")
            print(HyperLPR_plate_recognition(image))
            time.sleep(1)
        time.sleep(1)
except KeyboardInterrupt:
    GPIO.cleanup()
    cap.release()