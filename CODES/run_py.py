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
    result = test_dict.get(key)

    #return search_time

def display_key_names(test_dict):
    key_names = list(test_dict.keys())
   # print("Key Names:")
   # print("\n".join(key_names))

def main():
    dict_size = 1000
    search_times = [50, 500, 5000, 50000]

    test_dict = create_dict(dict_size)
    display_key_names(test_dict)

    for x in range(5):
        print("---------------------------")
        for search_time in search_times:
            start_time = time.time()
            for i in range(search_time):
                random_key = random.choice(list(test_dict.keys()))
                #time_taken = search_dict(random_key, test_dict)
                result = test_dict.get(random_key)
                #total_time += time_taken
            end_time = time.time()
            total_time = end_time - start_time             
            print(f"Search time for {search_time} searches:= { round(total_time * 1000, 1)} ")

if __name__ == '__main__':
    main()
