# coding: utf-8
import pprint
import re  # noqa: F401

import six
from enum import Enum
from schema.aws.autoscaling.ec2instancelaunchsuccessful.Details import Details  # noqa: F401,E501

class EC2InstanceLaunchSuccessful(object):


    _types = {
        'Details': 'Details',
        'Description': 'str',
        'EndTime': 'datetime',
        'RequestId': 'str',
        'ActivityId': 'str',
        'Cause': 'str',
        'AutoScalingGroupName': 'str',
        'StartTime': 'datetime',
        'EC2InstanceId': 'str',
        'StatusCode': 'str',
        'StatusMessage': 'str',
        'Origin': 'str',
        'Destination': 'str'
    }

    _attribute_map = {
        'Details': 'Details',
        'Description': 'Description',
        'EndTime': 'EndTime',
        'RequestId': 'RequestId',
        'ActivityId': 'ActivityId',
        'Cause': 'Cause',
        'AutoScalingGroupName': 'AutoScalingGroupName',
        'StartTime': 'StartTime',
        'EC2InstanceId': 'EC2InstanceId',
        'StatusCode': 'StatusCode',
        'StatusMessage': 'StatusMessage',
        'Origin': 'Origin',
        'Destination': 'Destination'
    }

    def __init__(self, Details=None, Description=None, EndTime=None, RequestId=None, ActivityId=None, Cause=None, AutoScalingGroupName=None, StartTime=None, EC2InstanceId=None, StatusCode=None, StatusMessage=None, Origin=None, Destination=None):  # noqa: E501
        self._Details = None
        self._Description = None
        self._EndTime = None
        self._RequestId = None
        self._ActivityId = None
        self._Cause = None
        self._AutoScalingGroupName = None
        self._StartTime = None
        self._EC2InstanceId = None
        self._StatusCode = None
        self._StatusMessage = None
        self._Origin = None
        self._Destination = None
        self.discriminator = None
        self.Details = Details
        self.Description = Description
        self.EndTime = EndTime
        self.RequestId = RequestId
        self.ActivityId = ActivityId
        self.Cause = Cause
        self.AutoScalingGroupName = AutoScalingGroupName
        self.StartTime = StartTime
        self.EC2InstanceId = EC2InstanceId
        self.StatusCode = StatusCode
        self.StatusMessage = StatusMessage
        self.Origin = Origin
        self.Destination = Destination


    @property
    def Details(self):

        return self._Details

    @Details.setter
    def Details(self, Details):


        self._Details = Details


    @property
    def Description(self):

        return self._Description

    @Description.setter
    def Description(self, Description):


        self._Description = Description


    @property
    def EndTime(self):

        return self._EndTime

    @EndTime.setter
    def EndTime(self, EndTime):


        self._EndTime = EndTime


    @property
    def RequestId(self):

        return self._RequestId

    @RequestId.setter
    def RequestId(self, RequestId):


        self._RequestId = RequestId


    @property
    def ActivityId(self):

        return self._ActivityId

    @ActivityId.setter
    def ActivityId(self, ActivityId):


        self._ActivityId = ActivityId


    @property
    def Cause(self):

        return self._Cause

    @Cause.setter
    def Cause(self, Cause):


        self._Cause = Cause


    @property
    def AutoScalingGroupName(self):

        return self._AutoScalingGroupName

    @AutoScalingGroupName.setter
    def AutoScalingGroupName(self, AutoScalingGroupName):


        self._AutoScalingGroupName = AutoScalingGroupName


    @property
    def StartTime(self):

        return self._StartTime

    @StartTime.setter
    def StartTime(self, StartTime):


        self._StartTime = StartTime


    @property
    def EC2InstanceId(self):

        return self._EC2InstanceId

    @EC2InstanceId.setter
    def EC2InstanceId(self, EC2InstanceId):


        self._EC2InstanceId = EC2InstanceId


    @property
    def StatusCode(self):

        return self._StatusCode

    @StatusCode.setter
    def StatusCode(self, StatusCode):


        self._StatusCode = StatusCode


    @property
    def StatusMessage(self):

        return self._StatusMessage

    @StatusMessage.setter
    def StatusMessage(self, StatusMessage):


        self._StatusMessage = StatusMessage


    @property
    def Origin(self):

        return self._Origin

    @Origin.setter
    def Origin(self, Origin):


        self._Origin = Origin


    @property
    def Destination(self):

        return self._Destination

    @Destination.setter
    def Destination(self, Destination):


        self._Destination = Destination

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
        if issubclass(EC2InstanceLaunchSuccessful, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        return self.to_str()

    def __eq__(self, other):
        if not isinstance(other, EC2InstanceLaunchSuccessful):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        return not self == other

