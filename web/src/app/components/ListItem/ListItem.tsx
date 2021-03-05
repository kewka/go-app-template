import clsx from 'clsx';
import { ComponentProps } from 'react';
import styles from './ListItem.module.scss';

type Props = ComponentProps<'li'> & {
  divider?: boolean;
};

export default function ListItem({
  divider,
  children,
  className,
  ...restProps
}: Props) {
  return (
    <li
      className={clsx(styles.root, divider && styles.divider, className)}
      {...restProps}
    >
      {children}
    </li>
  );
}
