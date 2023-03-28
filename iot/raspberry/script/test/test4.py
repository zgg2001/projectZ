#!/usr/bin/python3
# -*- coding: UTF-8 -*-

import RPi.GPIO as GPIO
import time
from hyperlpr import *
import cv2
import os

# 管道
pipe_read  = "../../tmp/pipe.2"
pipe_write = "../../tmp/pipe.1"
rf = os.open(pipe_read, os.O_RDONLY | os.O_NONBLOCK)
wf = os.open(pipe_write, os.O_SYNC | os.O_CREAT | os.O_RDWR)

while True:
    ret = os.read(rf, 10).decode()
    if len(ret) != 0:
        print("1", ret)
