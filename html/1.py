import time
import cv2  # pip install opencv-python -i 镜像源网址
from email.mime.image import MIMEImage  # 用来构造邮件内容的库
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart
import smtplib  # 发送邮件

def GetPicture():
    """
    拍照保存图像
    :return:
    """
    # 创建一个窗口
    cv2.namedWindow('camera', 1)
    # 调用摄像头   IP摄像头APP
    video = "http://admin:admin@172.19.48.1:8081/video"
    cap = cv2.VideoCapture(video)
    while True:
        success, img = cap.read()
        cv2.imshow("camera", img)
        # 按键处理
        key = cv2.waitKey(10)
        if key == 27:
            # esc
            break
        if key == 32:
            # 空格
            fileaname = 'frames.jpg'
            cv2.imwrite(fileaname, img)

    # 释放摄像头
    cap.release()
    # 关闭窗口
    cv2.destroyWindow("camera")
