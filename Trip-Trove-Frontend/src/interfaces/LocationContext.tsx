import { ReactNode } from "react";
import { ILocation } from "./Location";

export interface ILocationContext {
  locations: ILocation[];
  getLocationById: (id: number) => Promise<ILocation | undefined>;
}

export interface ILocationProviderProps {
  children: ReactNode;
}
