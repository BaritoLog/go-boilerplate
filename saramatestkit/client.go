package saramatestkit

import "github.com/Shopify/sarama"

type Client struct {
	ConfigFunc             func() *sarama.Config
	ControllerFunc         func() (*sarama.Broker, error)
	RefreshControllerFunc  func() (*sarama.Broker, error)
	BrokersFunc            func() []*sarama.Broker
	BrokerFunc             func(brokerID int32) (*sarama.Broker, error)
	TopicsFunc             func() ([]string, error)
	PartitionsFunc         func(topic string) ([]int32, error)
	WritablePartitionsFunc func(topic string) ([]int32, error)
	LeaderFunc             func(topic string, partitionID int32) (*sarama.Broker, error)
	ReplicasFunc           func(topic string, partitionID int32) ([]int32, error)
	InSyncReplicasFunc     func(topic string, partitionID int32) ([]int32, error)
	OfflineReplicasFunc    func(topic string, partitionID int32) ([]int32, error)
	RefreshBrokersFunc     func(addrs []string) error
	RefreshMetadataFunc    func(topics ...string) error
	GetOffsetFunc          func(topic string, partitionID int32, time int64) (int64, error)
	CoordinatorFunc        func(consumerGroup string) (*sarama.Broker, error)
	RefreshCoordinatorFunc func(consumerGroup string) error
	InitProducerIDFunc     func() (*sarama.InitProducerIDResponse, error)
	CloseFunc              func() error
	ClosedFunc             func() bool
}

func NewClient() *Client {

	return &Client{
		ConfigFunc:             func() *sarama.Config { return nil },
		ControllerFunc:         func() (*sarama.Broker, error) { return nil, nil },
		RefreshControllerFunc:  func() (*sarama.Broker, error) { return nil, nil },
		BrokersFunc:            func() []*sarama.Broker { return []*sarama.Broker{} },
		BrokerFunc:             func(brokerID int32) (*sarama.Broker, error) { return nil, nil },
		TopicsFunc:             func() ([]string, error) { return []string{}, nil },
		PartitionsFunc:         func(topic string) ([]int32, error) { return []int32{}, nil },
		WritablePartitionsFunc: func(topic string) ([]int32, error) { return []int32{}, nil },
		LeaderFunc:             func(topic string, partitionID int32) (*sarama.Broker, error) { return nil, nil },
		ReplicasFunc:           func(topic string, partitionID int32) ([]int32, error) { return []int32{}, nil },
		InSyncReplicasFunc:     func(topic string, partitionID int32) ([]int32, error) { return []int32{}, nil },
		OfflineReplicasFunc:    func(topic string, partitionID int32) ([]int32, error) { return nil, nil },
		RefreshBrokersFunc:     func(addrs []string) error { return nil },
		RefreshMetadataFunc:    func(topics ...string) error { return nil },
		GetOffsetFunc:          func(topic string, partitionID int32, time int64) (int64, error) { return 0, nil },
		CoordinatorFunc:        func(consumerGroup string) (*sarama.Broker, error) { return nil, nil },
		RefreshCoordinatorFunc: func(consumerGroup string) error { return nil },
		InitProducerIDFunc:     func() (*sarama.InitProducerIDResponse, error) { return nil, nil },
		CloseFunc:              func() error { return nil },
		ClosedFunc:             func() bool { return false },
	}
}

func (c *Client) Config() *sarama.Config                        { return c.ConfigFunc() }
func (c *Client) Controller() (*sarama.Broker, error)           { return c.ControllerFunc() }
func (c *Client) RefreshController() (*sarama.Broker, error)    { return c.RefreshControllerFunc() }
func (c *Client) Brokers() []*sarama.Broker                     { return c.BrokersFunc() }
func (c *Client) Broker(brokerID int32) (*sarama.Broker, error) { return c.BrokerFunc(brokerID) }
func (c *Client) Topics() ([]string, error)                     { return c.TopicsFunc() }
func (c *Client) Partitions(topic string) ([]int32, error)      { return c.PartitionsFunc(topic) }
func (c *Client) WritablePartitions(topic string) ([]int32, error) {
	return c.WritablePartitionsFunc(topic)
}
func (c *Client) Leader(topic string, partitionID int32) (*sarama.Broker, error) {
	return c.LeaderFunc(topic, partitionID)
}
func (c *Client) Replicas(topic string, partitionID int32) ([]int32, error) {
	return c.ReplicasFunc(topic, partitionID)
}
func (c *Client) InSyncReplicas(topic string, partitionID int32) ([]int32, error) {
	return c.InSyncReplicasFunc(topic, partitionID)
}
func (c *Client) OfflineReplicas(topic string, partitionID int32) ([]int32, error) {
	return c.OfflineReplicasFunc(topic, partitionID)
}
func (c *Client) RefreshBrokers(addrs []string) error    { return c.RefreshBrokersFunc(addrs) }
func (c *Client) RefreshMetadata(topics ...string) error { return c.RefreshMetadataFunc(topics...) }
func (c *Client) GetOffset(topic string, partitionID int32, time int64) (int64, error) {
	return c.GetOffsetFunc(topic, partitionID, time)
}
func (c *Client) Coordinator(consumerGroup string) (*sarama.Broker, error) {
	return c.CoordinatorFunc(consumerGroup)
}
func (c *Client) RefreshCoordinator(consumerGroup string) error {
	return c.RefreshCoordinatorFunc(consumerGroup)
}
func (c *Client) InitProducerID() (*sarama.InitProducerIDResponse, error) {
	return c.InitProducerIDFunc()
}
func (c *Client) Close() error { return c.CloseFunc() }
func (c *Client) Closed() bool { return c.ClosedFunc() }
