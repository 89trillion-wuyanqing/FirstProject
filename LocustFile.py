
from locust import HttpUser, TaskSet, task, between


class UserBehavior(TaskSet):

    @task(1)  # 任务的权重为1，如果有多个任务，可以将权重值定义成不同的值，
    def get_Armys(self):

        data = {
            "rarity": "2",
            "unlockArena": "2",
            "cvc" : "1000"
        }
        response = self.client.post('/getArmys', data = data,name="get_Armys")
        if not response.ok:
            print(response.text)
            response.failure('Got wrong response')

    @task(2)  # 任务的权重为1，如果有多个任务，可以将权重值定义成不同的值，
    def get_ArmyRarity(self):

        data2 = {
            "id": "10501"

        }

        response = self.client.post('/getArmyRarity', data = data2, name="get_ArmyRarity")
        if not response.ok:
            print(response.text)
            response.failure('Got wrong response')

    @task(3)  # 任务的权重为1，如果有多个任务，可以将权重值定义成不同的值，
    def get_ArmyAtk(self):


        data3 = {
            "id": "10502"

        }
        response = self.client.post('/getArmyAtk', data=data3,name="get_ArmyAtk")
        if not response.ok:
            print(response.text)
            response.failure('Got wrong response')

    @task(4)  # 任务的权重为1，如果有多个任务，可以将权重值定义成不同的值，
    def get_ArmysByCvc(self):

        data4 = {
            "cvc": "1000"

        }
        response = self.client.post('/getArmysByCvc', data=data4 , name="get_ArmysByCvc")
        if not response.ok:
            print(response.text)
            response.failure('Got wrong response')

    @task(5)  # 任务的权重为1，如果有多个任务，可以将权重值定义成不同的值，
    def get_ArmysByStage(self):

        response = self.client.get('/getArmysByStage', name="get_ArmysByStage")
        if not response.ok:
            print(response.text)
            response.failure('Got wrong response')



class TestLocust(HttpUser):

    tasks = [UserBehavior]
    wait_time = between(2, 5)
    host = "http://127.0.0.1:8000"
    #task_set = UserBehavior
    #host = "http://127.0.0.1/:8000"  # 被测服务器地址
    #min_wait = 5000
# 最小等待时间，即至少等待多少秒后Locust选择执行一个任务。

    #max_wait = 9000
# 最大等待时间，即至多等待多少秒后Locust选择执行一个任务。