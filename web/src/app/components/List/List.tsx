import clsx from 'clsx';
import { ComponentProps, ReactNode } from 'react';
import styles from './List.module.scss';

type Props = ComponentProps<'ul'> & {
  subheader?: ReactNode;
};
export default function List({
  subheader,
  children,
  className,
  ...restProps
}: Props) {
  return (
    <ul className={clsx(styles.root, className)} {...restProps}>
      {subheader && <div className={styles.subheader}>{subheader}</div>}
      {children}
    </ul>
  );
}
