import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import ResponsiveAppBar from "../components/ResponsiveAppBar";
import Footer from "../components/Footer";
import { Box, Typography, Divider } from "@mui/material";
import { IDestination } from "../interfaces/Destination";
import { ILocation } from "../interfaces/Location";
import { useDestinations } from "../contexts/DestinationContext";
import { useLocations } from "../contexts/LocationContext";

export function Destination() {
  const { id } = useParams<{ id: string }>();
  const { getDestinationById } = useDestinations();
  const { getLocationById } = useLocations();
  const [destination, setDestination] = useState<IDestination | undefined>();
  const [location, setLocation] = useState<ILocation | undefined>();
  const nav = useNavigate();

  useEffect(() => {
    const fetchDestinationDetails = async () => {
      if (!id) {
        return null;
      }
      const fetchedDestination = await getDestinationById(parseInt(id));
      setDestination(fetchedDestination);
    };
    fetchDestinationDetails();
    console.log(destination)

    const fetchLocationDetails = async () => {
      if (!destination) {
        return null;
      }
      if (!destination.location_id) {
        return null;
      }
      const fetchedLocation = await getLocationById(destination.location_id);
      setLocation(fetchedLocation);
    };
    fetchLocationDetails();
    console.log(location)

  });

  if (!destination || !location) {
    return (
      <Typography variant="h4" align="center">
        Loading...
      </Typography>
    );
  } else {
    return (
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "space-between",
          backgroundColor: "#f9f9f9",
          minHeight: "100vh",
          width: "100vw",
        }}
      >
        <ResponsiveAppBar />
        <Box
          sx={{
            borderRadius: "1em",
            overflow: "hidden",
            boxShadow: "0 10px 30px rgba(0,0,0,0.1)",
            width: "100%",
            maxWidth: "1080px",
          }}
        >
          <Box sx={{ width: "100%", overflow: "hidden" }}>
            <img
              src={destination.image_url}
              alt="Destination"
              style={{
                width: "100%",
                maxHeight: "500px",
                objectFit: "cover",
              }}
            />
          </Box>
          <Box
            sx={{
              p: 3,
              overflowY: "auto",
              maxHeight: "calc(100vh - 300px - 64px)",
            }}
          >
            <Typography
              variant="h4"
              component="h2"
              sx={{ textAlign: "center", mb: 2 }}
            >
              {destination.name}
            </Typography>
            <Divider sx={{ mb: 2 }} />
            <Typography variant="h6" component="h3" sx={{ mb: 1 }}>
              Location: {location.name}
            </Typography>
            <Typography variant="subtitle1" component="h3" sx={{ mb: 1 }}>
              Country: {location.country}
            </Typography>
            <Typography variant="body1" sx={{ mb: 2 }}>
              {destination.description}
            </Typography>
          </Box>
        </Box>
        <Footer />
      </Box>
    );
  }
}
