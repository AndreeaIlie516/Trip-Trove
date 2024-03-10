import { ReactNode } from "react";
import { IDestination } from "./Destination";

export interface IDestinationContext {
  destinations: IDestination[];
  getDestinationById: (id: number) => Promise<IDestination | undefined>;
  //addDestination: (destination: Destination) => void;
  //updateDestination: (destination: Destination) => void;
  //deleteDestination: (id: number) => void;
}

export interface IDestinationProviderProps {
  children: ReactNode;
}
