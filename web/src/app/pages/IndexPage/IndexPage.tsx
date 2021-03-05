import * as api from 'app/api';
import Button from 'app/components/Button/Button';
import Header from 'app/components/Header/Header';
import List from 'app/components/List/List';
import ListItem from 'app/components/ListItem/ListItem';
import Main from 'app/components/Main/Main';
import { useCallback, useEffect, useState } from 'react';
import uniqBy from 'lodash/uniqBy';

function useItems() {
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<api.schema['service.ItemsListResponse']>();
  const fetchMore = useCallback(async () => {
    setLoading(true);
    api
      .getItems({
        params: {
          offset: data?.items.length,
        },
      })
      .then(({ data }) => {
        setData((prevData) => ({
          count: data.count,
          items: uniqBy([...(prevData?.items || []), ...data.items], 'id'),
        }));
      })
      .finally(() => setLoading(false));
  }, [data]);
  useEffect(() => {
    setLoading(true);
    api
      .getItems()
      .then(({ data }) => setData(data))
      .finally(() => setLoading(false));
  }, []);
  return {
    data,
    loading,
    fetchMore,
  };
}

export default function IndexPage() {
  const { data, loading, fetchMore } = useItems();
  return (
    <>
      <Header title="App" />
      <Main>
        {data && (
          <List subheader={`Items (${data.count})`}>
            {data.items.map((item) => (
              <ListItem divider key={item.id}>
                {item.name}
              </ListItem>
            ))}
            {data.items.length < data.count && (
              <Button
                onClick={fetchMore}
                disabled={loading}
                block
                color="secondary"
                variant="contained"
              >
                Load more
              </Button>
            )}
          </List>
        )}
      </Main>
    </>
  );
}
