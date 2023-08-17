from selenium import webdriver
from selenium.webdriver.common.by import By
import time
driver = webdriver.Chrome()
driver.get("http://39.101.167.251/qftest")
driver.maximize_window()
time.sleep(2)
# driver.find_element(by=By.LINK_TEXT, value="免费注册").click()
# driver.find_element(by=By.ID, value="username").send_keys("hxpeee")
# driver.find_element(by=By.ID, value="email").send_keys("8912345@qq.com")
# driver.find_element(by=By.ID, value="password").send_keys("123456")
# driver.find_element(by=By.ID, value="repassword").send_keys("123456")
# driver.find_element(by=By.LINK_TEXT, value="立即注册").click()
# time.sleep(2)
# driver.find_element(by=By.LINK_TEXT, value="退出登录").click()

driver.find_element(by=By.LINK_TEXT, value="登陆").click()
driver.find_element(by=By.ID, value="username").send_keys("hxpeee")
driver.find_element(by=By.ID, value="password").send_keys("123456")
time.sleep(3)

driver.find_element(by=By.LINK_TEXT, value="登     陆").click()
time.sleep(3)
