import Nanobus from "nanobus"
import { Events } from "./events"

export const bus = new Nanobus<Events>()
