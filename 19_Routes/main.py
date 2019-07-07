
import multiprocessing
from urllib.request import urlopen

def f(url):
    req = urlopen(url)
    try:
        print(len(req.read()))
    finally:
        req.close()


urls = ("https://reactjs.org/", "https://python.org", "https://golang.org")


if __name__ == "__main__":
    p = multiprocessing.Pool(3)
    p.map(f, urls)

