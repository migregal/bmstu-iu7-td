import Placeholder from "components/Placeholder"
import { ReactComponent as Icon404 } from "./404.svg"
import { Link } from "react-router-dom"
import { PATH } from "routes/paths"
import Button from "components/Button"


export function NotFoundPage() {
  return <Placeholder
    shouldShow
    message={<>
      <h2>Can&apos;t find this page</h2>
      <Link to={PATH.INDEX}>
        <Button outline>Return to home page</Button>
      </Link>
    </>}
  >
    <Icon404 />
  </Placeholder>
}
