# coding: utf-8
import pprint
import re  # noqa: F401

import six
from enum import Enum

class Details(object):


    _types = {
        'Subnet_ID': 'str',
        'Availability_Zone': 'str'
    }

    _attribute_map = {
        'Subnet_ID': 'Subnet ID',
        'Availability_Zone': 'Availability Zone'
    }

    def __init__(self, Subnet_ID=None, Availability_Zone=None):  # noqa: E501
        self._Subnet_ID = None
        self._Availability_Zone = None
        self.discriminator = None
        self.Subnet_ID = Subnet_ID
        self.Availability_Zone = Availability_Zone


    @property
    def Subnet_ID(self):

        return self._Subnet_ID

    @Subnet_ID.setter
    def Subnet_ID(self, Subnet_ID):


        self._Subnet_ID = Subnet_ID


    @property
    def Availability_Zone(self):

        return self._Availability_Zone

    @Availability_Zone.setter
    def Availability_Zone(self, Availability_Zone):


        self._Availability_Zone = Availability_Zone

    def to_dict(self):
        result = {}

        for attr, _ in six.iteritems(self._types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(Details, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        return self.to_str()

    def __eq__(self, other):
        if not isinstance(other, Details):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        return not self == other

