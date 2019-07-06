import time
t0 = time.time()
time.sleep(3.5)
t1 = time.time()

print("Took {:.2f} seconds".format(t1-t0), "To Delay")