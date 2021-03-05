import clsx from 'clsx';
import { ComponentProps } from 'react';
import styles from './Main.module.scss';

type Props = ComponentProps<'main'> & {
  fullHeight?: boolean;
};

export default function Main({ fullHeight, ...restProps }: Props) {
  return (
    <main
      {...restProps}
      className={clsx(
        styles.root,
        fullHeight && styles.fullHeight,
        restProps.className
      )}
    />
  );
}
