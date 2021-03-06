import React from 'react'
import { Table, Skeleton } from 'antd'
import classNames from 'classnames'
import Card from '../Card'
import styles from './index.module.less'

class TableCard extends React.PureComponent {
  render() {
    const {
      title,
      className,
      style,
      loading,
      loadingSkeletonRows,
      cardExtra,
      ...rest
    } = this.props
    return (
      <Card
        title={title}
        style={style}
        className={classNames(styles.cardTable, className)}
        extra={cardExtra}
      >
        {loading ? (
          <Skeleton
            active
            title={false}
            paragraph={{ rows: loadingSkeletonRows || 5 }}
          />
        ) : (
          <div className={styles.cardTableContent}>
            <Table pagination={false} size="middle" {...rest} />
          </div>
        )}
      </Card>
    )
  }
}

export default TableCard
