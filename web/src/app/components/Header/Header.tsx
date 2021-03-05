import clsx from 'clsx';
import { ComponentProps } from 'react';
import styles from './Header.module.scss';

type Props = ComponentProps<'header'> & {
  title?: string;
};

export default function Header({ title, className, ...restProps }: Props) {
  return (
    <header className={clsx(styles.root, className)} {...restProps}>
      <div className={styles.container}>
        <div className={styles.title}>{title}</div>
      </div>
    </header>
  );
}
