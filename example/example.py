import json

# echo is a function that returns the input
def echo(input_str: str) -> str:
    return input_str

# add is a function that returns the sum of the number a and b in the
# input json string.
def add(input_str: str) -> str:
    input_dict = json.loads(input_str)
    output_dict = {"sum": input_dict["a"] + input_dict["b"]}
    return json.dumps(output_dict)


def print_numbers(nonsence: str):
    for i in range(10):
        print(i)
    return nonsence


if __name__ == "__main__":
    input_str = "[1,2,3]"
    print("echo({}) = {}".format(input_str, echo(input_str)))

    input_str = '{"a": 1, "b": 2}'
    print("add({}) = {}".format(input_str, add(input_str)))