/***************************************************
 * @Time : 2021/07/29 22:38 pm
 * @Author : Bai Neng
 * @File : rabbmitmq
 * @Software: GoLand
 **************************************************/
package RabbitMQ

import (
	"errors"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

var (
	ErrInvalidConfig = errors.New("invalid pool config")
	ErrPoolClosed    = errors.New("pool closed")
)

type PoolConfig struct {
	MaxOpen     int           // 池中最大资源数
	NumOpen     int           // 当前池中资源数
	MinOpen     int           // 池中最少资源数
	Closed      bool          // 池是否已关闭
	IdleTimeout time.Duration //空闲连接连接超时时间
	WaitTimeOut time.Duration //等待获取连接超时时间
}
type NewConnection func() (*amqp.Connection, error)
type RabbitmqPool struct {
	mu    sync.Mutex
	conns chan *amqp.Connection

	newConnection func() (*amqp.Connection, error)
	poolConfig    *PoolConfig
}

func NewPool(config *PoolConfig, newConnection NewConnection) (*RabbitmqPool, error) {
	if config.MaxOpen <= 0 || config.MinOpen > config.MaxOpen {
		return nil, ErrInvalidConfig
	}
	p := &RabbitmqPool{
		conns:         make(chan *amqp.Connection, config.MaxOpen),
		newConnection: newConnection,
		poolConfig:    config,
	}
	for i := 0; i < config.MinOpen; i++ {
		conn, err := newConnection()
		if err != nil {
			continue
		}
		config.NumOpen++
		p.conns <- conn
	}
	return p, nil
}

func (p *RabbitmqPool) Get() (*amqp.Connection, error) {
	if p.poolConfig.Closed {
		return nil, ErrPoolClosed
	}
	for {
		conn, err := p.connection()
		if err != nil {
			return nil, err
		}
		// todo maxLifttime处理
		return conn, nil
	}
}

func (p *RabbitmqPool) connection() (*amqp.Connection, error) {
	select {
	case conn := <-p.conns:
		return conn, nil
	default:
		p.mu.Lock()
		if p.poolConfig.NumOpen >= p.poolConfig.MaxOpen {
			conn := <-p.conns
			p.mu.Unlock()
			return conn, nil
		}
		// 新建连接
		conn, err := p.newConnection()
		if err != nil {
			p.mu.Unlock()
			return nil, err
		}
		p.poolConfig.NumOpen++
		p.mu.Unlock()
		return conn, nil
	}
}

// 释放单个资源到连接池
func (p *RabbitmqPool) Release(conn *amqp.Connection) error {
	if p.poolConfig.Closed {
		return ErrPoolClosed
	}
	p.mu.Lock()
	p.conns <- conn
	p.mu.Unlock()
	return nil
}

// 关闭单个资源
func (p *RabbitmqPool) Close(conn *amqp.Connection) error {
	p.mu.Lock()
	conn.Close()
	p.poolConfig.NumOpen--
	p.mu.Unlock()
	return nil
}

// 关闭连接池，释放所有资源
func (p *RabbitmqPool) ClosePool() error {
	if p.poolConfig.Closed {
		return ErrPoolClosed
	}
	p.mu.Lock()
	close(p.conns)
	for conn := range p.conns {
		conn.Close()
		p.poolConfig.NumOpen--
	}
	p.poolConfig.Closed = true
	p.mu.Unlock()
	return nil
}

//打开通道
func OpenChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	return ch, err
}
