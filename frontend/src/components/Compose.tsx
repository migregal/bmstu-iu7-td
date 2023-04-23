import { ComponentType, PropsWithChildren } from "react"

type Props = {
    components: ComponentType<PropsWithChildren>[]
}

export default function Compose({ components, children}: PropsWithChildren<Props>) {
  return <>
    {components.reduceRight(
      (children, Component) => <Component>{children}</Component>,
      children,
    )}
  </>
}
