import { ElementType } from 'react';

type Props = {
  component: ElementType;
} & Record<any, any>;

export default function Overridable({
  component: Component,
  ...restProps
}: Props) {
  return <Component {...restProps} />;
}
