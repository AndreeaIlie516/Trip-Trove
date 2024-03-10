import React from 'react'
import ReactDOM from 'react-dom/client'
import RoutesProvider from './router/router'
import { DestinationProvider } from './contexts/DestinationContext'
import { LocationProvider } from './contexts/LocationContext'

ReactDOM.createRoot(
  document.getElementById('root')!).render(
    <React.StrictMode>
      <DestinationProvider>
        <LocationProvider>
          <RoutesProvider />
        </LocationProvider>
      </DestinationProvider>
    </React.StrictMode>,
  )