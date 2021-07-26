from locust import HttpUser, TaskSet, task,between


class QuickstartUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def baidu_index1(self):
        self.client.get("/getLicitSoldier?rarity=1&unK=1&cvc=1000")

    @task
    def baidu_index2(self):
        self.client.get("/getRarity?Id=19809")

    @task
    def baidu_index3(self):
        self.client.get("/getCombatPoints?Id=19809")

    @task
    def baidu_index4(self):
        self.client.get("/getCvcLicitSoldier?cvc=1000")

    @task
    def baidu_index5(self):
        self.client.get("/getUnkSoldier")