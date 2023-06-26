import time
import random
import string

def create_dict(dict_size):
    test_dict = {}
    prefix = string.ascii_lowercase + "_"
    for i in range(dict_size):
        key = prefix + str(i).zfill(3)
        test_dict[key] = i
    return test_dict


def search_dict(key, test_dict):
    start_time = time.time_ns()
    result = test_dict.get(key)
    end_time = time.time_ns()
    search_time = end_time - start_time 
    return search_time

def display_key_names(test_dict):
    key_names = list(test_dict.keys())
   # print("Key Names:")
   # print("\n".join(key_names))

def main():
    dict_size = 1000
    search_times = [50, 500, 5000]

    test_dict = create_dict(dict_size)
    display_key_names(test_dict)

    for search_time in search_times:
        total_time = 0
        for i in range(search_time):
            random_key = random.choice(list(test_dict.keys()))
            time_taken = search_dict(random_key, test_dict)
            total_time += time_taken
        average_time = total_time / search_time
        print(f"\nAverage search time for {search_time} searches: {average_time:.2f} nanoseconds")

if __name__ == '__main__':
    main()
