import RPi.GPIO as GPIO
import time
from hyperlpr import *
import cv2
import numpy as np

# 摄像头超声波
Trig_Pin1 = 22
Echo_Pin1 = 27

GPIO.setmode(GPIO.BCM)
GPIO.setup(Trig_Pin1, GPIO.OUT, initial = GPIO.LOW)
GPIO.setup(Echo_Pin1, GPIO.IN)

# 读取摄像头设备/设置分辨率
cap = cv2.VideoCapture(0) #'/dev/video0', cv2.CAP_V4L
#cap.set(cv2.CAP_PROP_FRAME_WIDTH, 1280)
#cap.set(cv2.CAP_PROP_FRAME_HEIGHT, 720)
cap.set(3,640)
cap.set(4,480)
ret, frame = cap.read()
rows, cols, channels = frame.shape
print(cols, rows, channels)

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

while(1):
        ret,dst = cap.read()
        cv2.imshow('usb camera', dst)
        cv2.waitKey(1)
        #if (k == ord('q')):
        #    break
        #elif(k == ord('s')):
        #    #name = input('name:')
        #    name += 1
        #    #filename = r'./camera/' + str(name) + '.jpg'
        #    filename = str(name) + '.jpg'
        #    cv2.imwrite(filename, dst)
        #    print(filename)
        #    #break 
        print("helloworld")
cap.release()
cv2.destroyAllWindows()
