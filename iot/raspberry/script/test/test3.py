import cv2
import numpy as np
name = 0
cap = cv2.VideoCapture(0)
 
cap.set(3,640)
cap.set(4,480)
 
ret, frame = cap.read()
rows, cols, channels = frame.shape
print(cols, rows, channels)
 
# 图像预处理
def img_p(img):
 
    # 灰度化
    gray_img = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
 
    # 平滑滤波
    blur = cv2.blur(gray_img, (3,3))
 
    # 二值化
    ret1, th1 = cv2.threshold(blur, 190, 255, cv2.THRESH_BINARY)
 
    # 透视变换
    b = 50
    pts1 = np.float32([[b, 0], [cols-b, 0], [0, rows], [cols, rows]])
    pts2 = np.float32([[0, 0], [cols, 0], [0, rows], [cols, rows]])
    M = cv2.getPerspectiveTransform(pts1, pts2)
    dst = cv2.warpPerspective(blur, M, (cols, rows))
 
    return dst
 
 
while(1):
        ret,frame = cap.read()
        dst = img_p(frame)
        cv2.imshow('usb camera', dst)
 
        k = cv2.waitKey(1)
        if (k == ord('q')):
            break
        elif(k == ord('s')):
            #name = input('name:')
            name += 1
            #filename = r'./camera/' + str(name) + '.jpg'
            filename = str(name) + '.jpg'
            cv2.imwrite(filename, dst)
            print(filename)
            #break 
cap.release()
cv2.destroyAllWindows()