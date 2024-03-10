import { createContext, useContext, useEffect, useState } from "react";
import {
  ILocationContext,
  ILocationProviderProps,
} from "../interfaces/LocationContext";
import { ILocation } from "../interfaces/Location";
import baseUrl from "../consts";

const LocationContext = createContext<ILocationContext | undefined>(undefined);
const locationUrl = `${baseUrl}/locations`;

export function useLocations() {
  const context = useContext(LocationContext);
  if (!context) {
    throw new Error("no context location");
  }
  return context;
}

export const LocationProvider: React.FC<ILocationProviderProps> = ({
  children,
}) => {
  const [locations, setLocation] = useState<ILocation[]>([]);

  useEffect(() => {
    const fetchLocations = async () => {
      try {
        const response = await fetch(locationUrl);
        if (!response.ok) {
          throw new Error("Could not fetch locations");
        }
        const data = (await response.json()) as ILocation[];
        console.log("data", data);
        setLocation(data);
      } catch (error: unknown) {
        console.error(`Error`);
      }
    };
    fetchLocations();
  }, []);

  const getLocationById = async (
    id: number
  ): Promise<ILocation | undefined> => {
    try {
      const response = await fetch(`${locationUrl}/${id}`);
      if (!response.ok) {
        throw new Error("could not find location by id");
      }
      const data = (await response.json()) as ILocation;
      return data;
    } catch (err: unknown) {
      console.log(`error`);
    }
  };

  return (
    <LocationContext.Provider value={{ locations, getLocationById }}>
      {children}
    </LocationContext.Provider>
  );
};
