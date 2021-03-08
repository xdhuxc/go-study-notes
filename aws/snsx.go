package aws

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	log "github.com/sirupsen/logrus"
)

type SNSClient struct {
	sns   *sns.SNS
	topic string
}

func NewSNSClient(region string, disableSSL bool) (*SNSClient, error) {
	if region == "" {
		sess, err := session.NewSession()
		if err != nil {
			return nil, err
		}
		metadata := ec2metadata.New(sess)
		hostRegion, err := metadata.Region()
		if err != nil {
			return nil, fmt.Errorf("the AWS service requires a region.")
		}
		region = hostRegion
	}
	var c *aws.Config

	AccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	SecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if AccessKey != "" && SecretKey != "" { // 使用环境变量中的 AccessKey 和 SecretKey
		creds := credentials.NewStaticCredentials(AccessKey, SecretKey, "")
		c = &aws.Config{
			Region:      aws.String(region),
			Credentials: creds,
			DisableSSL:  aws.Bool(disableSSL),
		}
	} else {
		c = &aws.Config{ // 使用 ~/.aws/credentials 文件中的 AccessKey 和 SecretKey 。
			Region:     aws.String(region),
			DisableSSL: aws.Bool(disableSSL),
		}
	}

	sess, err := session.NewSession(c)
	if err != nil {
		return nil, err
	}
	snsClient := sns.New(sess)

	return &SNSClient{
		sns: snsClient,
	}, nil
}

func (snsx *SNSClient) SetTopic(topic string) {
	snsx.topic = topic
}

func (snsx *SNSClient) ListSubscriptions() error {
	var token *string
	subscriptionInput := sns.ListSubscriptionsInput{
		NextToken: aws.String(""),
	}

	for {
		result, err := snsx.sns.ListSubscriptions(&subscriptionInput)
		if err != nil {
			if awsError, ok := err.(awserr.Error); ok {
				switch awsError.Code() {
				case sns.ErrCodeInvalidParameterException:
					log.Errorln(sns.ErrCodeInvalidParameterException, awsError.Error())
				case sns.ErrCodeInternalErrorException:
					log.Errorln(sns.ErrCodeInternalErrorException, awsError.Error())
				case sns.ErrCodeNotFoundException:
					log.Errorln(sns.ErrCodeNotFoundException, awsError.Error())
				case sns.ErrCodeAuthorizationErrorException:
					log.Errorln(sns.ErrCodeAuthorizationErrorException, awsError.Error())
				default:
					log.Errorln(awsError.Error())
				}
			} else {
				log.Errorln(err)
			}

			return err
		}

		subscriptions := result.Subscriptions
		/**
		由于主题已经被删除，不能根据主题获取订阅，只能循环所有的订阅，删除指定主题的订阅
		*/
		for _, subscription := range subscriptions {
			if *subscription.TopicArn == snsx.topic {
				fmt.Println(*subscription.TopicArn, "--->", *subscription.SubscriptionArn)
				log.Infof("deleting subscriptionArn %s", subscription.String())
			}
		}

		token = result.NextToken
		if token != nil {
			subscriptionInput.SetNextToken(*token)
		} else {
			break
		}
	}

	return nil
}

func (snsx *SNSClient) listSubscriptions() error {
	var token *string
	var total int64
	subscriptionInput := sns.ListSubscriptionsInput{
		NextToken: aws.String(""),
	}

	for {
		result, err := snsx.sns.ListSubscriptions(&subscriptionInput)
		if err != nil {
			if awsError, ok := err.(awserr.Error); ok {
				switch awsError.Code() {
				case sns.ErrCodeInvalidParameterException:
					log.Errorln(sns.ErrCodeInvalidParameterException, awsError.Error())
				case sns.ErrCodeInternalErrorException:
					log.Errorln(sns.ErrCodeInternalErrorException, awsError.Error())
				case sns.ErrCodeNotFoundException:
					log.Errorln(sns.ErrCodeNotFoundException, awsError.Error())
				case sns.ErrCodeAuthorizationErrorException:
					log.Errorln(sns.ErrCodeAuthorizationErrorException, awsError.Error())
				default:
					log.Errorln(awsError.Error())
				}
			} else {
				log.Errorln(err)
			}

			return err
		}

		subscriptions := result.Subscriptions
		/**
		由于主题已经被删除，不能根据主题获取订阅，只能循环所有的订阅，删除指定主题的订阅
		*/
		var wg sync.WaitGroup
		for _, subscription := range subscriptions {
			if *subscription.TopicArn == snsx.topic {
				total = total + 1
				fmt.Println(*subscription.TopicArn, "--->", *subscription.SubscriptionArn)
				log.Infof("deleting subscriptionArn %s", subscription.GoString())
				wg.Add(1)
				go func(s *SNSClient, subscriptionArn string) {
					defer wg.Done()
					s.deleteSubscription(subscriptionArn)
				}(snsx, *subscription.SubscriptionArn)
			}
		}
		wg.Wait()

		token = result.NextToken
		if token != nil {
			subscriptionInput.SetNextToken(*token)
		} else {
			break
		}
	}

	return nil
}

func (snsx *SNSClient) deleteSubscription(subscriptionArn string) {
	unsubscribeInput := sns.UnsubscribeInput{
		SubscriptionArn: &subscriptionArn,
	}

	_, err := snsx.sns.Unsubscribe(&unsubscribeInput)
	if err != nil {
		log.Errorln(err)
	}
}

func (snsx *SNSClient) DeleteSubscription(subscriptionArn string) error {
	unsubscribeInput := sns.UnsubscribeInput{
		SubscriptionArn: &subscriptionArn,
	}

	_, err := snsx.sns.Unsubscribe(&unsubscribeInput)
	if err != nil {
		return err
	}

	return nil
}
