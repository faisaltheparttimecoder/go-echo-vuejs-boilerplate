# Generic
echo "# Setting up environment variables"
export LOG_LEVEL="DEBUG"
export API_PORT=8082

# Cleanup function
cleanup ()
{
    echo "# Cleaning up the previous go process [if any]"
    lsof -i 4tcp:${API_PORT} -sTCP:LISTEN | grep -v "COMMAND" | awk '{print $2}' | xargs -n1 /bin/kill
}

# Run cleanup before starting
cleanup

# If only backend is needed
echo "# Starting on the setup"
npm --prefix ./web run build
go run *.go

# Cleanup when using ctrl + C
trap cleanup EXIT