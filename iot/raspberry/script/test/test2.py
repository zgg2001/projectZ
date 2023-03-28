import cv2 as cv

# 读取设备
cap = cv.VideoCapture('/dev/video0', cv.CAP_V4L)

# set dimensions 设置分辨率
cap.set(cv.CAP_PROP_FRAME_WIDTH, 800)
cap.set(cv.CAP_PROP_FRAME_HEIGHT, 400)

ret, frame = cap.read()
if ret:
    cv.imwrite('./tmp/tmp.jpg', frame) # 截图

cap.release()