from locust import HttpUser, task, between
import random
import string


class MyUser(HttpUser):
    wait_time = between(1, 3)  # Time between consecutive tasks
    nonce = None  # Variable to store the nonce for pq request
    server_nonce = None  # Variable to store the server_nonce for dh request
    pq_message_id = 1  # Initial message_id for pq request

    def generate_random_nonce(self):
        # Generate a random 20-character string
        return ''.join(random.choices(string.ascii_letters + string.digits, k=20))

    def generate_random_positive_int(self):
        # Generate a random positive integer
        return random.randint(1, 1000)
    
    def generate_random_message_id(self):
        # Generate a random odd number for message_id
        return random.randint(1, 1000) * 2 + 1


    @task
    def pq_handler(self):
        # Generate a new nonce for pq request
        nonce = self.generate_random_nonce()
        msg_id = self.generate_random_message_id()
        # Make the pq request with the generated nonce and odd message_id
        pq_request = {
            "messageId": msg_id,
            "nonce": nonce
        }

        print(pq_request)
        response = self.client.post("/auth/pq", json=pq_request)

        self.nonce = nonce
        self.pq_message_id = msg_id
        # Extract the server_nonce from the response
        self.server_nonce = response.json().get("nonces", {}).get("nonce_server")

    @task
    def dh_handler(self):
        if self.server_nonce is None or self.nonce is None:
            return  # Skip if pq_handler hasn't run or failed

        # Increment message_id for dh request, ensuring it's greater than pq_message_id
        dh_message_id = self.pq_message_id + 2

        # Generate a random positive integer for A
        A_value = self.generate_random_positive_int()

        # Pair the pq_nonce and server_nonce to send in the dh request
        dh_request = {
            "messageId": dh_message_id,
            "nonces": {
                "nonce": self.nonce,
                "nonceServer": self.server_nonce
            },
            "A": str(A_value)
        }
        self.client.post("/auth/dh", json=dh_request)
