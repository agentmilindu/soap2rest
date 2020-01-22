from jinja2 import Environment, FileSystemLoader
import os
import yaml


root = os.path.dirname(os.path.abspath(__file__))
templates_dir = os.path.join(root, 'templates')
env = Environment( loader = FileSystemLoader(templates_dir) )
template = env.get_template('main.template.go')


filename = os.path.join(root, 'main.go')
config = {}

with open("config.yaml", 'r') as stream:
    try:
        config = yaml.safe_load(stream)
    except yaml.YAMLError as e:
        print(e)

print(config)

with open(filename, 'w') as fh:
    fh.write(template.render(
        config = config
    ))
