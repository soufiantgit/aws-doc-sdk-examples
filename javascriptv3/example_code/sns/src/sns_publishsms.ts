/* Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0

ABOUT THIS NODE.JS EXAMPLE: This example works with AWS SDK for JavaScript version 3 (v3),
which is available at https://github.com/aws/aws-sdk-js-v3. This example is in the 'AWS SDK for JavaScript v3 Developer Guide' at
https://docs.aws.amazon.com/sdk-for-javascript/v3/developer-guide//sns-examples-sending-sms.html.

Purpose:
sns_publishsms.ts demonstrates how to use Amazon SNS to send an SMS text message to a phone number.

Inputs (replace in code):
- REGION
- TEXT_MESSAGE
- PHONE_NUMBER

Running the code:
ts-node sns_publishsms.ts
 */
// snippet-start:[sns.JavaScript.SMS.publishV3]
// Import required AWS SDK clients and commands for Node.js
const { SNSClient, PublishCommand } = require("@aws-sdk/client-sns");

// Set the AWS Region
const REGION = "region"; //e.g. "us-east-1"

// Set the parameters
const params = {
  Message: "MESSAGE_TEXT" /* required */,
  PhoneNumber: "PHONE_NUMBER", //PHONE_NUMBER, in the E.164 phone number structure
};

// Create SNS service object
const sns = new SNSClient({ region: REGION });

const run = async () => {
  try {
    const data = await sns.send(new PublishCommand(params));
    console.log("Success, message published. MessageID is " + data.MessageId);
  } catch (err) {
    console.error(err, err.stack);
  }
};
run();
// snippet-end:[sns.JavaScript.SMS.publishV3]

