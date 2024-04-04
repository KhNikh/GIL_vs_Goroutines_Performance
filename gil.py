import threading
import time

# A CPU-bound function that increments a counter a large number of times
def cpu_bound_task(i):
    counter = 0
    for _ in range(10**7):
        counter += 1
    

# Function to run the CPU-bound task in multiple threads
def run_tasks_in_threads(num_threads):
    threads = []
    for i in range(num_threads):
        thread = threading.Thread(target=cpu_bound_task, args=(i, ))
        thread.start()
        threads.append(thread)
    
    for thread in threads:
        thread.join()

if __name__ == "__main__":
    start_time = time.time()
    run_tasks_in_threads(10)  # Run the CPU-bound task in 10 threads
    end_time = time.time()
    print("Execution time:", end_time - start_time, "seconds")
