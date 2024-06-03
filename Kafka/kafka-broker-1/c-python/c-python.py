from kafka import KafkaProducer
import json
import time

# Địa chỉ IP của Kafka broker, thay bằng địa chỉ IP của máy bạn
broker_ip = '192.168.1.61'
broker_port = '9092'
broker_address = f'{broker_ip}:{broker_port}'

# Khởi tạo Kafka producer
producer = KafkaProducer(
    bootstrap_servers=[broker_address],
    value_serializer=lambda v: json.dumps(v).encode('utf-8')  # Chuyển dữ liệu thành JSON
)

# Tên topic
topic_name = 'message-log'

# Dữ liệu mẫu để gửi
data = {
    'vector': [0.1, 0.2, 0.3, 0.4]  # Ví dụ về vector embedding
}

try:
    while True:
        # Gửi dữ liệu tới Kafka broker
        producer.send(topic_name, value=data)
        print(f'Sent data: {data}')
        time.sleep(1)  # Gửi mỗi giây
except KeyboardInterrupt:
    print("Stopped by user")
finally:
    # Đóng kết nối Kafka producer
    producer.close()