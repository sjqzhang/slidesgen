#! -*- encoding: utf-8 -*-

from setuptools import setup, find_packages

setup(
    name='slidesgen',
    version='0.1',
    description='A simple tool to generate slides from markdown files',
    author='sjqzhang',
    packages=find_packages(),
    include_package_data=True,
    install_requires=[
        'click',
    ],
    entry_points='''
        [console_scripts]
        slidesgen=slidesgen:main
            
    ''',
)
