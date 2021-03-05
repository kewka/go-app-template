import clsx from 'clsx';
import { ComponentProps } from 'react';
import Overridable from '../Overridable/Overridable';
import styles from './Button.module.scss';

type Props = ComponentProps<'button'> & {
  variant?: 'text' | 'outlined' | 'contained';
  color?: 'primary' | 'secondary' | 'default';
  block?: boolean;
} & Partial<ComponentProps<typeof Overridable>>;

const classByVariant: Record<NonNullable<Props['variant']>, string> = {
  text: styles.variantText,
  outlined: styles.variantOutlined,
  contained: styles.variantContained,
};

const classByColor: Record<NonNullable<Props['color']>, string> = {
  primary: styles.colorPrimary,
  secondary: styles.colorSecondary,
  default: styles.colorDefault,
};

export default function Button({
  variant = 'text',
  color = 'default',
  block,
  component = 'button',
  className,
  ...restProps
}: Props) {
  return (
    <Overridable
      className={clsx(
        styles.root,
        classByVariant[variant],
        classByColor[color],
        block && styles.block,
        className
      )}
      component={component}
      {...restProps}
    />
  );
}
